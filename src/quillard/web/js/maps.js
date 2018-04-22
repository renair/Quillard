let map = undefined;
let personageMarker = undefined;
let homeMarker = undefined;

function initMap() {
	const options = {
		zoom: 16,
		streetViewControl: false,
		disableDefaultUI:true,
		scaleControl: true,
		gestureHandling: "cooperative",
		center: {lat: 50.433072, lng: 30.62024}
	};
	map = new google.maps.Map(document.getElementById('googlemap'), options);
}

function setHomeMarker(position) {
	const pos = toLatLng(position);
	if(!homeMarker){
		homeMarker = new google.maps.Marker({
			position: pos,
			map: map,
			title: position.name,
			icon:"/icons/home.png"
		});
	} else {
		homeMarker.setPosition(pos);
	}
	map.panTo(pos);
}

function addHomeMarker() {
	
}

function setPersonageMarker(personage) {
	const pos = toLatLng(personage.position);
	if(!personageMarker){
		personageMarker = new google.maps.Marker({
			position: pos,
			map: map,
			title: personage.name,
			icon:"/icons/personage.png",
			draggable: true
		});
	} else {
		personageMarker.setPosition(pos);
		personageMarker.setTitle(personage.name);
	}
	map.panTo(pos);
}

//function required on registration

let mapClickedListener = undefined;

function mapOnClicked(evt) {
	const coords = {
		lat: evt.latLng.lat(),
		lng: evt.latLng.lng()
	}
	setHomeMarker(fromLatLng(coords));
	$("#register_form > input[name=longitude]").val(evt.latLng.lng());
	$("#register_form > input[name=latitude]").val(evt.latLng.lat());
}

function setMapClicker() {
	mapClickedListener = map.addListener("click", mapOnClicked);
}

function disableMapClicker() {
	google.maps.event.removeListener(mapClickedListener);
}

//helper functions

function toLatLng(position) {
	return {
		lat: position.latitude,
		lng: position.longitude
	}
}

function fromLatLng(pos) {
	return {
		latitude: pos.lat,
		longitude: pos.lng
	}
}

//closure for moving personage marker
function getPersonageMover(personage) {
	return function(){
		setPersonageMarker(personage);
	}
}

//closure for moving map
function getMapMover(position) {
	return function() {
		const pos = toLatLng(position);
		map.panTo(pos);
	}
}