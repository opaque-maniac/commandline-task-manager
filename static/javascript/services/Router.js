const routes = {
    "/": "home-page",
    "/edit": "edit-page",

};

const Router = {
    init: () => {
        document.querySelectorAll("a.navlink").forEach(a => {
            a.addEventListener("click", event => {
                event.preventDefault();
                const path = event.currentTarget.dataset.path;
                Router.go(path);
            });
        });

        window.addEventListener("popstate", event => {
            Router.go(event.state?.path || "/", false);
        });
        const path = location.pathname

        if (path === "/") {
            const main = document.querySelector("main");
            if (!main.classList.contains("main-home")) {
                main.classList.add("main-home");
            }
        }

        Router.go(path, false);
    },

    go: (url, addToHistory = true) => {
        if (addToHistory) {
            window.history.pushState({ path: url }, "", url);
        }

        const main = document.querySelector("main");
        if (main.classList.contains("main-home")) {
            if (url !== "/") {
                main.classList.remove("main-home");
            }
        } else {
            if (url === "/") {
                main.classList.add("main-home");
            }
        }
        

        const path = url.includes("/edit") ? "/edit" : url;
        const pageElement = document.createElement(routes[path] || "not-found-page");

        document.querySelector("main").replaceChildren(pageElement);

        window.scrollTo(0, 0);
    }
};

export default Router;

