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
  const due_Table = document.querySelector('.approved-data');
  var dueCount = 1;
  dues.forEach(due => {
    if (user.Department === due.Department && due.Status === "Approved") {
      console.log(due)
      due_Table.innerHTML += `
              
            <tr>
            <td>${dueCount}</td>
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
      dueCount += 1;
    }
  });

  if (dueCount === 1) {
    due_Table.innerHTML = `
      <tr>
      <td colspan="5">
          No Request Approved
      </td>
  </  tr>
      `;
  }
}