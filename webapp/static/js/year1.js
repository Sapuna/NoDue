window.onload = () => {
    fetch("/nodue")
        .then(response => response.text())
        .then(data => {
            const parsedData = JSON.parse(data);
            DisplayDue(parsedData);
        });
}

function DisplayDue(data){
    const noDueTable = document.querySelector('.no-due-data');
    var slno = 1;
    data.forEach(nodue => {
        if(nodue.Year === 1){
            noDueTable.innerHTML += `
            <tr>
            <td>${slno}</td>
            <td>${nodue.EnrollmentNumber}</td>
            <td>${nodue.Name}</td>
            <td>${nodue.Year}</td>
            <td>
            <span class="action_btn">
            <a href="#" style="background-color: #C8E7C8;">${nodue.Status}</a>
            </span>
            </td>
            </tr>
            
            `
        }
    });
}

// function DisplayDue(data) {
//     const groupedData = {};
//     var approved = 0;
//     data.forEach(due => {
//         if (due.Status === "Approved") {
//             const enrollmentNumber = due.EnrollmentNumber;
//             if (!groupedData[enrollmentNumber]) {
//                 groupedData[enrollmentNumber] = [];
//             }
//             groupedData[enrollmentNumber].push(due);
//             approved +=1;
//         }
//     });
    
//     console.log(approved)
//     for (const enrollmentNumber in groupedData) {
//         if (Object.hasOwnProperty.call(groupedData, enrollmentNumber)) {
//             const data = groupedData[enrollmentNumber];
//             console.log(data);
//             approved += 1;
//         }
//     }
//     console.log(approved)
// }
