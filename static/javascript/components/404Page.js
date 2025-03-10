import Router from "../services/Router.js";

export default class NotFoundPage extends HTMLElement {
    constructor() {
        super();
        this.root = this.attachShadow({mode: "open"})
        const styles = document.createElement("style");

        async function loadCSS() {
            const response = await fetch("/static/stylesheets/404Page.css");
            const content =  await response.text();
            styles.textContent = content;
        }
        loadCSS();
        this.root.appendChild(styles);
    }
    connectedCallback() {
        const template = document.getElementById("404-page-template");
        const content = template.content.cloneNode(true)

        // links
        this.root
            .querySelectorAll("a.navlink")
            .forEach(route => {
                Router.go(route.getAttribute("href"));
            });

        this.root.appendChild(content)
    }
}

customElements.define("not-found-page", NotFoundPage)
