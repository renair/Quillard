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
			title: position.name
		});
	} else {
		homeMarker.setPosition(pos);
	}
	map.panTo(pos);
}

function setPersonageMarker(personage) {
	const pos = toLatLng(personage.position);
	if(!personageMarker){
		personageMarker = new google.maps.Marker({
			position: pos,
			map: map,
			title: personage.name,
			draggable: true
		});
	} else {
		personageMarker.setPosition(pos);
	}
	map.panTo(pos);
}

function toLatLng(position) {
	return {
		lat: position.latitude,
		lng: position.longitude
	}
}

//closure for moving map
function getPersonageMover(personage) {
	return function(){
		setPersonageMarker(personage);
	}
}