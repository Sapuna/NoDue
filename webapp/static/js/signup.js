const form = document.getElementById('signup-form');
const nameInput = document.getElementById('name');
const emailInput = document.getElementById('email');
const userIdInput = document.getElementById('User-id');
const passwordInput = document.getElementById('password');

const confirmPasswordInput = document.getElementById('confirmPassword');

function setErrorFor(input, message) {
  const formField = input.parentElement;
  formField.classList.remove("success");
  formField.classList.add("error");
  const error = formField.querySelector("small");

  const inputFiled = input;
  inputFiled.classList.add("error-input")
  inputFiled.classList.remove("success-input")

  error.textContent = message;
}

function setSuccessFor(input) {
  const formField = input.parentElement;
  formField.classList.remove("error");
  formField.classList.add("success");

  const inputFiled = input;
  inputFiled.classList.remove("error-input")
  inputFiled.classList.add("success-input")

  const success = formField.querySelector("small");

  success.textContent = ' ';

}


//the form get validated when user click the sign up button
form.addEventListener('submit', function (event) {
  event.preventDefault();
  checkName();
  checkEmail();
  checkUserId();
  checkPasswordSame();
});

function checkName() {
  const nameValue = nameInput.value.trim();

  if (nameValue === '') {
    setErrorFor(nameInput, 'Name cannot be blank');
  } else if (!isValidName(nameValue)) {
    setErrorFor(nameInput, 'Name must contain only letters');
  } else {
    setSuccessFor(nameInput);
  }
}

function isValidName(name) {
  return /^[a-zA-Z]+$/.test(name);
}

function checkEmail() {
  const emailValue = emailInput.value.trim();

  if (emailValue === '') {
    setErrorFor(emailInput, 'Email cannot be blank');
  } else if (!isValidEmail(emailValue)) {
    setErrorFor(emailInput, 'Please enter a valid email address');
  } else {
    setSuccessFor(emailInput);
  }
}

function isValidEmail(email) {
  return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
}

function checkUserId() {
  const userIdValue = userIdInput.value.trim();
  const mini = 8;

  if (userIdValue === '') {
    setErrorFor(userIdInput, 'User ID cannot be blank');
  } else if (!isValidUserId(userIdValue)) {
    setErrorFor(userIdInput, 'User ID must be a number');
  } else if (userIdInput.value.length < mini) {
    setErrorFor(userIdInput, `User ID shoud have atleast ${mini} characters. For Example(12220089)`);
  } else {
    setSuccessFor(userIdInput);
  }
}

function isValidUserId(userId) {
  return /^\d+$/.test(userId);
}

function checkPasswordSame() {
  if (passwordInput.value !== confirmPasswordInput.value) {
    setErrorFor(confirmPasswordInput, "Password does not match")
  } else {
    setSuccessFor(confirmPasswordInput)
  }
}

//password strength test


passwordInput.addEventListener('input', function () {
  const passwordStrength = checkPasswordStrength(passwordInput.value);
  const passwordStrengthElement = document.getElementById('password-strength');
  passwordStrengthElement.textContent = passwordStrength;

});

function checkPasswordStrength(password) {

  let strength = 0;

  if (password.length < 6) {
    return 'Weak';
  }

  if (password.length >= 6 && password.length <= 8) {
    strength += 1;
  } else if (password.length > 8 && password.length <= 10) {
    strength += 2;
  } else if (password.length > 10) {
    strength += 3;
  }

  if (password.match(/[a-z]+/)) {
    strength += 1;
  }

  if (password.match(/[A-Z]+/)) {
    strength += 1;
  }

  if (password.match(/[0-9]+/)) {
    strength += 1;
  }

  if (password.match(/[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]+/)) {
    strength += 1;
  }

  if (strength <= 2) {
    return 'Weak';
  } else if (strength <= 4) {
    return 'Medium';
  } else {
    return 'Strong';
  }
}
