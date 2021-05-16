var form = document.forms["formConnect"];
var champLogin = document.forms["formConnect"]["login"];
var champPassword = document.forms["formConnect"]["password"];

function validateForm()
 {  
  if(champLogin.value === "")
   {
    champLogin.classList.remove("correct");
    champLogin.classList.add("incorrect");
    return false;
  }
  else if(champPassword.value === "")
   {
    champPassword.classList.remove("correct");
    champPassword.classList.add("incorrect");
    return false;
  }
                 
  else
   {
    return true;
  }
}

form.addEventListener("blur", function(event)
 {
  if(event.target.value === "")
   {
    event.target.classList.remove("correct");
    event.target.classList.add("incorrect");
  }
  else
   {
    event.target.classList.remove("incorrect");
    event.target.classList.add("correct");
  }
}, true);

erreurConnect.addEventListener("click", function() {
  erreurConnect.style.display="none";
});
