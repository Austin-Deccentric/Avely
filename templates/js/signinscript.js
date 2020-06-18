let x = document.getElementsByClassName("pwordeye")[0]; // hide password (eye slash)
let y = document.getElementsByClassName("pwordeye")[1]; // show password (eye)
let p = document.getElementById("userpword");
x.setAttribute('style', "visibility:hidden;")
// Toggle password visibility
x.addEventListener('click', myFunction);
y.addEventListener('click', myFunction);
function myFunction() {

    if (p.type === "password") {
        p.type = "text";
        y.setAttribute('style', "visibility:hidden;")
        x.setAttribute('style', "visibility:visible;")
    } else {
        p.type = "password";
        y.setAttribute('style', "visibility:visible;")
        x.setAttribute('style', "visibility:hidden;")
    }
}