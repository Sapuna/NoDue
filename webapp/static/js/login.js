const login_Form = document.querySelector("#login-form")

login_Form.addEventListener("submit", (e) => {
    e.preventDefault()
    Login();
})

const Login = () => {

    const userID = document.querySelector("#userID").value;
    const password = document.querySelector("#password").value

    var data = {
        UserID: parseInt(userID),
        Password: password
    }

    fetch('/login', {
        method: "POST",
        body: JSON.stringify(data),
        headers: { "Content-type": "application/json; charset=UTF-8" }
    })
        .then(response => {
            if (response.ok) {
                return response.json(); // Parse response data as JSON
            } else {
                throw new Error(response.statusText);
            }
        })
        .then(data => {
            if (data.error === "Unauthorized") {
                alert("Unauthorized. Credentials do not match!");
            } else {
                const user = data.data
                console.log(user)
                if (user.Type === "Student" && user.Department === "STUDENTS") {
                    localStorage.setItem("userID", user.ID);
                    window.open("./user/home.html", "_self")
                }
                else if (user.Type === "Department" && (user.Department === "LIBRARY" || user.Department === "FINANCE SECTION" || user.Department === "STUDENT SERVICE OFFICER" || user.Department === "ICT LAB" || user.Department === "COLLEGE CANTEEN" || user.Department === "ITERIA")) {
                    localStorage.setItem("userID", user.ID);
                    window.open("./dept/dept.html", "_self")
                }
                else if (user.Type === "Counselor" && user.Department === "ADMIN") {
                    localStorage.setItem("userID", user.ID);
                    window.open("./admin/vapp.html", "_self")

                }
            }
        });



};