const signUpform = document.querySelector('#signup-form')

signUpform.addEventListener('submit', (e) => {
    e.preventDefault()
    SignUP();
})

function SignUP() {

    const nameInput = document.getElementById('name');
    const emailInput = document.getElementById('email');
    const userIdInput = document.getElementById('User-id');
    const passwordInput = document.getElementById('password');
    const department = document.getElementById('Course');

    var data = {
        Name: nameInput.value,
        Email: emailInput.value,
        UserID: parseInt(userIdInput.value),
        Course: "Default course",
        Year: 10,
        Semester: "Default sem",
        Department: department.value,
        Password: passwordInput.value,
        Type: "Department"
    }

    console.log(data)

    fetch('/signup', {
        method: 'POST',
        body: JSON.stringify(data),
        headers: { 'Content-type': 'application/json; charset=UTF-8' }
    })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            alert("Signup Success!")
            window.open("../index.html", "_self")

        })
        .catch(error => {
            console.error(error);
        });
}