export async function fetchTodos() {
    try {
        const url = "/api/data";
        const options = {
            method: "GET",
            headers: {
                "Content-Type": "application/json"
            },
        };

        const response = await fetch(url, options);

        if (!response.ok) {
            throw new Error("Error fetching tasks");
        }

        const data = await response.json();
        return data.data ?? [];
    } catch(e) {
        throw e;
    }
}

