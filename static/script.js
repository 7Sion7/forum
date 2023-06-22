// script.js

function checkUsername(username) {
  var errorElement = document.getElementById("username-error");
  if (username === '') {
    errorElement.style.display = "none";
  }
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/check-username?username=" + encodeURIComponent(username), true);
  xhr.onreadystatechange = function () {
    if (xhr.readyState === XMLHttpRequest.DONE) {
      if (xhr.status === 200) {
        var response = JSON.parse(xhr.responseText);
        if (response.available) {
          errorElement.textContent = "";
          errorElement.style.display = "none"; // Clear any previous error message
        } else {
          errorElement.style.display = "block";
          errorElement.textContent = "Error:Username Already Exists";
        }
      } else {
        console.error("Error:", xhr.status);
      }
    }
  };
  xhr.send();
}
function checkEmail(email) {
  var errorElement = document.getElementById("username-error");
  if (email === '') {
    errorElement.style.display = "none";
  }
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/check-email?email=" + encodeURIComponent(email), true);
  xhr.onreadystatechange = function () {
    if (xhr.readyState === XMLHttpRequest.DONE) {
      if (xhr.status === 200) {
        var response = JSON.parse(xhr.responseText);
        if (response.available) {
          errorElement.textContent = "";
          errorElement.style.display = "none";
          // Clear any previous error message
        } else {
          errorElement.style.display = "block";
          errorElement.textContent = "Error:Email Already In Use";
        } 
      } else {
        console.error("Error:", xhr.status);
      }
    }
  };
  xhr.send();
}



// document.getElementById('signupForm').addEventListener('submit', function(e) {
//     /*This function should check with Go as to whether or not the input is valid through the DB Check*/
//     e.preventDefault(); // Prevent the form from submitting normally

//     var userId = 123; // Replace with the actual user ID

//     // Build the URL with the user ID and set it as the form's action
//     this.action = 'https://example.com/user/' + userId;

//     // Submit the form
//     this.submit();
// });

// document.getElementById('loginForm').addEventListener('submit', function(e) {
//     /*This function should check with Go as to whether or not the input is valid through the DB Check*/
//     e.preventDefault(); // Prevent the form from submitting normally

//     var userId = 123; // The ID will come from the Go backend

//     // Build the URL with the user ID and set it as the form's action
//     this.action = 'https://example.com/user/' + userId;

//     // Submit the form
//     this.submit();
// });


//NavBar
function IconBar(){
  var iconBar = document.getElementById("iconBar");
  var navigation = document.getElementById("navigation");
  if (navigation.classList.contains("hide")) {
    iconBar.setAttribute("style", "display:none;");
    navigation.classList.remove("hide");
  } else {
    iconBar.setAttribute("style", "display:block;");
    navigation.classList.add("hide");
  }
}

function incrementLikes(postID) {
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/likes?email=" + encodeURIComponent(postID), true);
    var likeCountElement = document.getElementById("likeCount");
    var likeCount = parseInt(likeCountElement.innerHTML);
    likeCount++;
    likeCountElement.innerHTML = likeCount;
}
function incrementDislikes(postID) {
    var dislikeCountElement = document.getElementById("dislikeCount");
    var dislikeCount = parseInt(dislikeCountElement.innerHTML);
    dislikeCount++;
    dislikeCountElement.innerHTML = dislikeCount;
}
function showComment(){
    var commentArea = document.getElementById("comment-area");
    if (commentArea.classList.contains("hide")) {
        commentArea.classList.remove("hide");
    } else {
        commentArea.classList.add("hide");
    }
    
}
