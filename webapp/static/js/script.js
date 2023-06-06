const header = document.querySelector("header");

window.addEventListener ("scroll", function() {
	header.classList.toggle ("sticky", window.scrollY > 0);
});

let menu = document.querySelector('#menu-icon');
let navbar = document.querySelector('.navbar');

menu.onclick = () => {
	menu.classList.toggle('bx-x');
	navbar.classList.toggle('open');
};

window.onscroll = () => {
	menu.classList.remove('bx-x');
	navbar.classList.remove('open');
};

// //uploading profile image
// let input_file = document.querySelector('input[type="file"]');
// let profile_img = document.querySelector('.profileadd label');

// input_file.onchange = (e) => {
// 	// accessing 1st choosen file
// 	let file = e.target.files[0];

// 	// Store "file's data" with the help of "URL Object"
// 	let url = URL.createObjectURL(file);

// 	// Changing background of 'label' and hiding user emoji icon
// 	profile_img.style.background = `url(${url}) center / cover no-repeat`;
// 	profile_img.querySelector('.user').style.display = "none";

// 	// Free up memory space (better perfomance)
// 	setTimeout(() => {
// 		URL.revokeObjectURL(url);
// 	}, 100)
// }



