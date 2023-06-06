const applyBtn = document.querySelector(".applybtn")


applyBtn.addEventListener("click", () => {
    ApplyDue()
})


const ApplyDue = async () => {
    console.log(localStorage.getItem("userID"));
    var data;
    try {
        const response = await fetch(`/users/${localStorage.getItem("userID")}`);
        data = await response.json();
    } catch (error) {
        console.error(error);
    }

    const Departments = ['LIBRARY', 'FINANCE SECTION', 'STUDENT SERVICE OFFICER', 'ICT LAB', 'COLLEGE CANTEEN', 'ITERIA'];
    console.log(data);
    var due = 0;
    Departments.forEach(department => {
        const _dueData = {
            EnrollmentNumber: data.UserID,
            Name: data.Name,
            Year: data.Year,
            Status: 'Pending',
            Department: department
        };

        fetch('/due', {
            method: 'POST',
            body: JSON.stringify(_dueData),
            headers: { 'Content-type': 'application/json; charset=UTF-8' }
        })
            .then(response => response.json())
            .then(data => {
                due += 1;
                console.log(data);
                // Handle the response data here
            })
            .catch(error => {
                console.error(error);
                // Handle any errors that occurred during the request
            });
    });

    if(due === 6){
        alert("due Added Successfully")
    }
}