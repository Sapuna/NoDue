window.onload = () => {
    fetch("/due")
        .then(response => response.text())
        .then(data => {
            const parsedData = JSON.parse(data);
            DisplayDue(parsedData);
        });
}

var DUEID = [];
async function DisplayDue(dues) {
    console.log(localStorage.getItem("userID"));
    var user;
    try {
        const response = await fetch(`/users/${localStorage.getItem("userID")}`);
        user = await response.json();
    } catch (error) {
        console.error(error);
    }
    console.log(user)
    const due_Table = document.querySelector('.dept-due-data')
    var dueCount = 1;
    dues.forEach(due => {
        if (user.Department === due.Department && due.Status === "Pending") {
            console.log(due);
            due_Table.innerHTML += `
            <tr>
                <td>${dueCount}</td>
                <td>${due.EnrollmentNumber}</td>
                <td>${due.Name}</td>
                <td>${due.Year}</td>
                <td>
                    <div class="container">
                        <div class="col-md-offset-5 col-md-5 check-row">
                            <form role="form">
                                <div class="form-group">
                                    <div class="checkbox">
                                        <label>
                                            <input type="checkbox" class="check" data-due-id="${due.ID}">
                                        </label>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </div>
                </td>
            </tr>`;

            dueCount += 1;
        }
    });

    if (dueCount === 1) {
        due_Table.innerHTML = `
        <tr>
        <td colspan="5">
            No Request
        </td>
    </  tr>
        `;
    }

    const checkboxes = document.querySelectorAll('.check');

    checkboxes.forEach(checkbox => {
        checkbox.addEventListener('change', function () {
            const dueId = this.getAttribute('data-due-id');
            if (this.checked) {
                DUEID.push(dueId);
            } else {
                const index = DUEID.indexOf(dueId);
                if (index !== -1) {
                    DUEID.splice(index, 1);
                }
            }
            console.log(DUEID);
        });
    });
}

const approveBtn = document.querySelector(".approve-btn")

approveBtn.addEventListener("click", () => {
    ApproveDue();
})


function ApproveDue() {
    DUEID.forEach((id) => {
        const approveData = {
            Status: 'Approved',
        };

        fetch(`/due/${id}`, {
            method: 'PUT',
            body: JSON.stringify(approveData),
            headers: { 'Content-type': 'application/json; charset=UTF-8' }
        })
            .then(response => response.json())
            .then(data => {
                alert(data)
                window.location.reload()
            }) 
            .catch(error => {
                console.error(error);
                alert("Error Approving")
            });
    })
}

const decline_btn = document.querySelector('.decline-btn')
decline_btn.addEventListener('click', () => {
    DeclineDue();
})

function DeclineDue() {
    DUEID.forEach((id) => {
        const approveData = {
            Status: 'Declined',
        };

        fetch(`/due/${id}`, {
            method: 'PUT',
            body: JSON.stringify(approveData),
            headers: { 'Content-type': 'application/json; charset=UTF-8' }
        })
            .then(response => response.json())
            .then(data => {
                alert(data)
                window.location.reload()
            })
            .catch(error => {
                console.error(error);
                alert("Error Approving")
            });
    })
}