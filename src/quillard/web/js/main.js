const LOGIN_URL = "/account/login";

let sessionKey = "";

function statusedAnswer(data){
	if(data.id != 0){
		alert(data.message)
	}
}

function onLogin(data){
	console.log(data)
	if(data.key && data.expires){
		console.log(data.key + " expire at " + data.expires);
		sessionKey = data.key;
		$("#login_form").hide();
		$("#welcome").show();
	} else {
		statusedAnswer(data);
	}
}

function renderPersonage() {
	
}

function createPersonageClicked() {
	
}

function loginClicked() {
	const email = $("#login_form input[name=email]").val();
	const password = $("#login_form input[name=password]").val();
	console.log(email + " " + password);
	const data = {
		"email": email,
		"password": password
	};
	$.post(LOGIN_URL, JSON.stringify(data), onLogin, "json");
}


$(document).ready(function(){
	$("#account_login_button").click(loginClicked);
});