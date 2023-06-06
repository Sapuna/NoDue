window.onload = () => {
    fetch("/due")
        .then(response => response.text())
        .then(data => {
            const parsedData = JSON.parse(data);
            DisplayDue(parsedData);
        });
}

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
    const decline_Table = document.querySelector('.declined-data');
    var declinedCount = 1;
    dues.forEach(due => {
        if (user.Department === due.Department && due.Status === "Declined") {
            decline_Table.innerHTML += `
            <tr>
            <td>${declinedCount}</td>
            <td>${due.EnrollmentNumber}</td>
            <td>${due.Name}</td>
            <td>${due.Year}</td>
			  <td>
				<span class="action_btn">
				  <a href="#" style="background-color: #C8E7C8;">${due.Status}</a>
				</span>
			  </td>
			</tr>
            `
            declinedCount += 1;
        }
    });
    if (declinedCount === 1) {
        decline_Table.innerHTML = `
        <tr>
        <td colspan="5">
            No Request Declined
        </td>
    </  tr>
        `;
    }
}