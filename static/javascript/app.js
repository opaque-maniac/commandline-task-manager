// pages
import HomePage from "./components/HomePage.js"
import NotFoundPage from "./components/404Page.js"
import EditPage from "./components/EditPage.js"

// todo item
import TodoItem from "./components/TodoItem.js"

// router
import Router from "./services/Router.js"

document.addEventListener("DOMContentLoaded", () => {
    Router.init();
    window.router = Router;
})
