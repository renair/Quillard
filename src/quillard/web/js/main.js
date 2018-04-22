//backend endpoints
const LOGIN_URL = "/account/login";
const REGISTER_URL = "/account/register";
const PERSONAGES_LIST = "/personage/list";
const PERSONAGE_CREATE = "/personage/create";
const HOME_POSITION = "/position/home";
const MY_RESOURCES = "/resources/myresources"

//templates
const PERSONAGE_CARD = $("#personage_card").html();
const RESOURCE = $("#resource").html();

//elements
const $pageHeader = $("#page_header");
const $resources = $pageHeader.find(".resources");
const $loginForm = $("#login_form");
const $registerForm = $("#register_form");
const $personages = $("#personages");
const $personageCreation = $("#personage_creation_form");

let session = {
	sessionKey: "",
	expirationTime: 0,
	personageId: 0,
	personagePosition: {},
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

function renderPersonageResourses(resources_data) {
	$resources.html("");
	console.log(resources_data);
	for(let i = 0; i < resources_data.length;i++) {
		const $res = $(RESOURCE);
		const res = resources_data[i];
		$res.text(res.name + " x" + res.amount);
		$res.attr("res-id", res.id);
		$resources.append($res);
	}
}

function personageChoseClicked() {
	$pageHeader.hide();
	$personages.show()
}

function getPersonageSelector(personage) {
	return function() {
		makePost(MY_RESOURCES, {personage_id:personage.id}, renderPersonageResourses);
		$pageHeader.find(".personage_name").text(personage.name);
		$pageHeader.find(".personage_name").off("click");
		$pageHeader.find(".personage_name").click(getMapMover(personage.position));
		$pageHeader.find(".home_location").off("click");
		$pageHeader.find(".home_location").click(() => {makePost(HOME_POSITION, {}, setHomeMarker);});
		$pageHeader.find(".chose_personage").click(personageChoseClicked);
		$pageHeader.show();
		$personages.hide();
	}
}

function renderPersonages(data){
	if(data.constructor == Array){
		$("#personages div[added=true]").remove();
		if(data.length >= 3){
			$("#personages > .personage_creating_panel").hide();
		} else {
			$("#personages > .personage_creating_panel").show();
		}
		for(let i = 0; i < data.length;++i){
			const personage = data[i];
			console.log(personage);
			let $card = $(PERSONAGE_CARD);
			$card.find(".personage_name").text(personage.name);
			$card.find(".personage_name").click(getPersonageSelector(personage));
			$card.find(".personage_position").text(personage.position.name);
			$card.find(".personage_position").click(getPersonageMover(personage));
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
		makePost(HOME_POSITION, {}, setHomeMarker);
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
	let userSession = JSON.parse(localStorage.getItem("userSession"));
	if(userSession && userSession.expirationTime > getUnixTime()){
		session = userSession;
		$loginForm.hide();
		makePost(PERSONAGES_LIST, {}, renderPersonages);
		makePost(HOME_POSITION, {}, setHomeMarker);
	} else {
		localStorage.removeItem("userSession");
	}
	$("#account_login_button").click(loginClicked);
	$("#open_register_button").click(openRegisterClicked);
	$("#account_register_button").click(registerClicked);
	$("#create_personage_button").click(createPersonageClicked);
});