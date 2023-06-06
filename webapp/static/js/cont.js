window.onload = () => {
    fetch("/due")
        .then(response => response.text())
        .then(data => {
            const parsedData = JSON.parse(data);
            DisplayDue(parsedData);
        });
}
var approveCount = 0;

async function DisplayDue(data) {

    console.log(localStorage.getItem("userID"));
    var user;
    try {
        const response = await fetch(`/users/${localStorage.getItem("userID")}`);
        user = await response.json();
    } catch (error) {
        console.error(error);
    }
    var datacount = 1;
    const due_Table = document.querySelector(".due-data")
    data.forEach(due => {
        if (user.UserID === due.EnrollmentNumber) {
            due_Table.innerHTML += `
            <tr>
            <td>${datacount}</td>
            <td>${due.Department}</td>
            <td>
            <span class="action_btn">
            <a href="#">${due.Status}</a>
            </span>
            </td>
            </tr>
            `
            datacount += 1;
        }
        if (due.Status === "Approved") {
            approveCount += 1;
        }
    });

    if (approveCount === 6) {
        document.querySelector('.submit-due-btn').style.opacity = "1";
        document.querySelector('.submit-due-btn').style.pointerEvents = "none";
        document.querySelector('.submit-due-btn').style.zindex = "2";
    }

}

const submitDuebtn = document.querySelector('.submit-due-btn')

submitDuebtn.addEventListener('click', () => {
    SubmitNoDue()
})

async function SubmitNoDue() {
    var user;

    try {
        const response = await fetch(`/users/${localStorage.getItem("userID")}`);
        user = await response.json();
    } catch (error) {
        console.error(error);
    }

    var data = {
        "EnrollmentNumber": user.UserID,
        "Name": user.Name,
        "Year": user.Year,
        "Status": "Approved"
    }

    fetch('/nodue', {
        method: 'POST',
        body: JSON.stringify(data),
        headers: { 'Content-type': 'application/json; charset=UTF-8' }
    })
        .then(response => response.json())
        .then(data => {

            alert(data.status)
        })
        .catch(error => {
            console.error(error);
        });
}