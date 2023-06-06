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
        if(nodue.Year === 2){
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
