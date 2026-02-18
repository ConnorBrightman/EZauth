const API_BASE = "";

function setToken(token) {
    localStorage.setItem("token", token);
}

function getToken() {
    return localStorage.getItem("token");
}

function logout() {
    localStorage.removeItem("token");
    window.location.href = "/login.html";
}

// LOGIN
async function loginUser(event) {
    event.preventDefault();

    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    const res = await fetch(API_BASE + "/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password })
    });

    const data = await res.json();

    if (res.ok) {
        setToken(data.token);
        window.location.href = "/me.html";
    } else {
        document.getElementById("message").innerText = data.error || "Login failed";
    }
}

// REGISTER
async function registerUser(event) {
    event.preventDefault();

    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    const res = await fetch(API_BASE + "/auth/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password })
    });

    const data = await res.json();

    if (res.ok) {
        window.location.href = "/login.html";
    } else {
        document.getElementById("message").innerText = data.error || "Register failed";
    }
}

// GET ME
async function loadMe() {
    const token = getToken();

    if (!token) {
        window.location.href = "/login.html";
        return;
    }

    const res = await fetch(API_BASE + "/auth/me", {
        headers: {
            "Authorization": "Bearer " + token
        }
    });

    if (!res.ok) {
        logout();
        return;
    }

    const data = await res.json();
    document.getElementById("user-info").innerText =
        JSON.stringify(data, null, 2);
}
