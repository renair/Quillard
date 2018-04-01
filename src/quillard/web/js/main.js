//backend endpoints
const LOGIN_URL = "/account/login";
const PERSONAGES_LIST = "/personage/list";
const PERSONAGE_CREATE = "/personage/create";

//templates
const PERSONAGE_CARD = $("#personage_card").html();

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
			$card.find(".personage_position").text(personage.position.name + "(" + personage.position.longitude + ";" + personage.position.latitude + ")");
			$("#personages").append($card);
		}
	} else {
		statusedAnswer(data);
	}
	$("#personages").show();
	$("#personage_creation_form").hide();
}

function onLogin(data){
	console.log(data)
	if(data.key && data.expires){
		const expireIn = data.expires - getUnixTime();
		console.log(data.key + " expire in " + expireIn);
		session.sessionKey = data.key;
		session.expirationTime = data.expires;
		localStorage.setItem("userSession", JSON.stringify(session));
		$("#login_form").hide();
		makePost(PERSONAGES_LIST, {}, renderPersonages);
	} else {
		statusedAnswer(data);
	}
}

function createPersonageClicked(){
	const $form = $("#personage_creation_form");
	$form.find("button").click(() => {
		data = {
			"name": $form.find(".personage_name").val()
		};
		makePost(PERSONAGE_CREATE, data, (response) => {
			makePost(PERSONAGES_LIST, {}, renderPersonages);
			statusedAnswer(response);
		});
	});
	$("#personages").hide();
	$form.show();
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


function getUnixTime() {
	return Math.floor(Date.now() / 1000);
}

$(document).ready(function(){
	initMap();
	setHomeMarker({latitude: 50.433077, longitude: 30.620238});
	let userSession = JSON.parse(localStorage.getItem("userSession"));
	if(userSession && userSession.expirationTime > getUnixTime()){
		session = userSession;
		$("#login_form").hide();
		makePost(PERSONAGES_LIST, {}, renderPersonages);
	} else {
		localStorage.removeItem("userSession");
	}
	$("#account_login_button").click(loginClicked);
	$("#create_personage_button").click(createPersonageClicked);
});