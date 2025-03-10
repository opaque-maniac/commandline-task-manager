import { completeTodo, deleteTodo } from "./Mutations.js";

function changeHandler(event, task) {
    event.preventDefault();
    completeTodo(task)
        .then(() => {
            event.currentTarget.checked = !event.currentTarget.checked;
        })
        .catch(error => {
            console.log(error);
        });
}

function deleteClickHandler(event, task) {
    event.preventDefault();
    deleteTodo(task)
        .then(() => {
            document.dispatchEvent(
                new CustomEvent("todo-reload")
            );
        })
        .catch(error => {
            console.log(error);
        });
}

function editClickHandler(event, task) {
    event.preventDefault();
    if (window.router) {
        window.router.go(`/edit?task=${encodeURI(task)}`);
    }
}

export default class TodoItem extends HTMLElement {
    constructor(task) {
        super();
        const template = document.getElementById("todo-item-template");
        const content = template.content.cloneNode(true);

        // update content
        content.querySelector("p").textContent = task.task;
        const input = content.querySelector("input")
        input.checked = task.completed;

        const callback = () => {
            content.classList.add("hidden");
        }

        // event listeners
        input
            .addEventListener("change", event => {
                changeHandler(event, task.task)
            });
        content
            .querySelector(".todo-update-button")
            .addEventListener("click",event => {
                editClickHandler(event, task.task)
            });
        content
            .querySelector(".todo-delete-button")
            .addEventListener("click", event => {
                deleteClickHandler(event, task.task, callback);
            });

        this.appendChild(content);
    }

    connectedCallback() {
    }
}

customElements.define("todo-item", TodoItem);
