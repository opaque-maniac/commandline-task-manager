import Router from "../services/Router.js";
import {updateTask} from "./Mutations.js";

export default class EditPage extends HTMLElement {
    constructor() {
        super();
        this.root = this.attachShadow({mode: "open"});
        const styles = document.createElement("style");

        async function loadCSS() {
            const response = await fetch("/static/stylesheets/EditPage.css");
            const content =  await response.text();
            styles.textContent = content;
        }
        loadCSS();
        this.root.appendChild(styles);
    }

    connectedCallback() {
        const template = document.getElementById("edit-page-template");
        const content = template.content.cloneNode(true);
        this.root.appendChild(content);

        let value = new URLSearchParams(window.location.search).get("task");

        if (!value) {
            Router.go("/404");
            return;
        }

        const input = this.root.querySelector(".update-todo-form input")
        input.value = value;

        const form = this.root.querySelector("form");
        form.addEventListener("submit", event => {
            event.preventDefault();
            const formData = new FormData(event.currentTarget)
            const todo = formData.get("todo");
            updateTask(value, todo)
                .then((data) => {
                    console.log(data);
                    Router.go("/");
                })
                .catch(error => {
                    console.log("Error", error.message);
                });
        });
    }

    async render() {
    }
}

customElements.define("edit-page", EditPage);
