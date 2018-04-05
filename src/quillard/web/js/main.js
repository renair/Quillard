//backend endpoints
const LOGIN_URL = "/account/login";
const REGISTER_URL = "/account/register";
const PERSONAGES_LIST = "/personage/list";
const PERSONAGE_CREATE = "/personage/create";

//templates
const PERSONAGE_CARD = $("#personage_card").html();

//elements
const $loginForm = $("#login_form");
const $registerForm = $("#register_form");
const $personages = $("#personages");
const $personageCreation = $("#personage_creation_form");

let session = {
	sessionKey: "",
	expirationTime: 0,
	personageId: 0,
	personagePosition: {}
}

//helper function
function makePost(url, data, successCallback){
	$.ajax({
		"type": "POST",
		"url": url,
		"success": successCallback,
		"data": JSON.stringify(data),
		"headers":{
			"Q-Session": session.sessionKey
		},
		"dataType": "json"
	});
}

function statusedAnswer(data){
	if(data.id != 0){
		alert(data.message)
	}
}

function renderPersonages(data){
	if(data.constructor == Array){
		$("#personages div[added=true]").remove();
		for(let i = 0; i < data.length;++i){
			personage = data[i];
			console.log(personage);
			let $card = $(PERSONAGE_CARD);
			$card.find(".personage_name").text(personage.name);
			$card.find(".personage_name").click(getPersonageMover(personage));
			$card.find(".personage_position").text(personage.position.name + "(" +
				 personage.position.longitude + ";" + personage.position.latitude + ")");
			$personages.append($card);
		}	
	} else {
		statusedAnswer(data);
	}
	$personages.show();
}

function onLogin(data){
	console.log(data)
	if(data.key && data.expires){
		const expireIn = data.expires - getUnixTime();
		console.log(data.key + " expire in " + expireIn);
		session.sessionKey = data.key;
		session.expirationTime = data.expires;
		localStorage.setItem("userSession", JSON.stringify(session));
		$loginForm.hide();
		$registerForm.hide();
		makePost(PERSONAGES_LIST, {}, renderPersonages);
	} else {
		statusedAnswer(data);
	}
}

function createPersonageClicked(){
	$personageCreation.find("button").click(() => {
		data = {
			"name": $personageCreation.find(".personage_name").val()
		};
		makePost(PERSONAGE_CREATE, data, (response) => {
			$personageCreation.hide();
			makePost(PERSONAGES_LIST, {}, renderPersonages);
			statusedAnswer(response);
		});
	});
	$personages.hide();
	$personageCreation.show();
}

function loginClicked() {
	const email = $loginForm.find("input[name=email]").val();
	const password = $loginForm.find("input[name=password]").val();
	console.log(email + " " + password);
	const data = {
		email: email,
		password: password
	};
	makePost(LOGIN_URL, data, onLogin);
}

function openRegisterClicked() {
	$loginForm.hide();
	$registerForm.show();
	setMapClicker();
}

function registerClicked() {
	const longitude = parseFloat($registerForm.find("input[name=longitude]").val());
	const latitude = parseFloat($registerForm.find("input[name=latitude]").val());
	const data = {
		email: $registerForm.find("input[name=email]").val(),
		password: $registerForm.find("input[name=password]").val(),
		home_position: {
			longitude: longitude,
			latitude: latitude,
			name: $registerForm.find("input[name=home_name]").val()
		}
	};
	makePost(REGISTER_URL, data, onLogin);
	disableMapClicker();
}


function getUnixTime() {
	return Math.floor(Date.now() / 1000);
}

$(document).ready(function(){
	initMap();
	setHomeMarker({latitude: 50.433077, longitude: 30.620238});
	let userSession = JSON.parse(localStorage.getItem("userSession"));
	if(userSession && userSession.expirationTime > getUnixTime()){
		session = userSession;
		$loginForm.hide();
		makePost(PERSONAGES_LIST, {}, renderPersonages);
	} else {
		localStorage.removeItem("userSession");
	}
	$("#account_login_button").click(loginClicked);
	$("#open_register_button").click(openRegisterClicked);
	$("#account_register_button").click(registerClicked);
	$("#create_personage_button").click(createPersonageClicked);
});