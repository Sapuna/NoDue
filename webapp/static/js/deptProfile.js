window.onload = () => {
    fetch(`/users/${localStorage.getItem("userID")}`)
        .then(response => response.text())
        .then(data => {
            const parsedData = JSON.parse(data);
            DisplayUser(parsedData);
        });
}



function DisplayUser(user) {

    const profileData = document.querySelector('.table_responsive')
    profileData.innerHTML = `

    <table>
				<thead>
					<tr>
						<th>Name</th>
						<th>${user.Name}</th>
					</tr>
				</thead>

				<thead>
					<tr>
						<th>Email Address</th>
						<th>${user.Email}</th>
					</tr>
				</thead>

				<thead>
					<tr>
						<th>Country</th>
						<th>Bhutan</th>
					</tr>
				</thead>

				</tbody>
			</table>
    
    `
}