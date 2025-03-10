import Router from "../services/Router.js";
import TodoItem from "./TodoItem.js";
import {fetchTodos} from "./Queries.js"
import { removeAll, sendNewTodo } from "./Mutations.js";

export default class HomePage extends HTMLElement {
    constructor() {
        super();
        this.root = this.attachShadow({mode: "open"});
        const styles = document.createElement("style");

        async function loadCSS() {
            const response = await fetch("/static/stylesheets/HomePage.css");
            const content =  await response.text();
            styles.textContent = content;
        }
        loadCSS();
        this.root.appendChild(styles);
    }

    connectedCallback() {
        const template = document.getElementById("home-page-template");
        const content = template.content.cloneNode(true);
        this.root.appendChild(content);

        // form
        this.root
            .querySelector(".new-todo-form")
            .addEventListener("submit", event => {
                event.preventDefault();
                const formData = new FormData(event.currentTarget);
                const todo = formData.get("todo");
                event.target.reset()
                sendNewTodo(todo)
                    .catch(error => {
                        console.log("Error", error.message);
                    })
                    .finally(() => {
                        document.dispatchEvent(new CustomEvent("todo-reload"));
                    });
            });

        // clear button
        this.root
            .querySelector(".clear-all-button")
            .addEventListener("click", event => {
                event.preventDefault();
                removeAll()
                    .then(() => {
                        document.dispatchEvent(new CustomEvent("todo-reload"));
                    })
                    .catch(error => {
                        console.log("Error", error.message);
                    });
            })

        window.addEventListener("reload-list", event => {
            event.preventDefault();
            this.render();
        })

        document.addEventListener("todo-reload", async () => {
            await this.render();
        });

        // render dynamic data
        this.render();
    }

    async render() {
        const todos = await fetchTodos();
        const section = this.root.querySelector(".todo-list");
        section.innerHTML = "";
        
        if (todos.length === 0) {
            if (!section.classList.contains("hidden")) {
                section.classList.add("hidden");
            }
            section.innerHTML = `
                    <h2 class="not-found">No todo tasks found</h2>`;
        } else {
            if (section.classList.contains("hidden")) {
                section.classList.remove("hidden");
            }
            const list = document.createElement("ul");
            todos.forEach(todo => {
                const listItem = document.createElement("li");
                const todoItem = new TodoItem(todo);;
                listItem.append(todoItem);
                list.appendChild(listItem);
            });
            section.appendChild(list);
        }
    }
}

customElements.define("home-page", HomePage);
