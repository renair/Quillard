//backend endpoints
const LOGIN_URL = "/account/login";
const PERSONAGES_LIST = "/personage/list";
const PERSONAGE_CREATE = "/personage/create";

//templates
const PERSONAGE_CARD = $("#personage_card").html();

let sessionKey = "";
let personageId = 0;
let personagePosition = {};

//helper function
function makePost(url, data, successCallback){
	$.ajax({
		"type": "POST",
		"url": url,
		"success": successCallback,
		"data": JSON.stringify(data),
		"headers":{
			"Q-Session": sessionKey
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
			$card.find(".personage_position").text("(" + personage.position.longitude + ";" + personage.position.latitude + ")");
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
		console.log(data.key + " expire at " + data.expires);
		sessionKey = data.key;
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


$(document).ready(function(){
	$("#account_login_button").click(loginClicked);
	$("#create_personage_button").click(createPersonageClicked);
});