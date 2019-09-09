function locateUser(map) {
	if (navigator.geolocation) {
		map.locate({setView: true });
	}
}

let searchCtrl = L.control.fuseSearch({panelTitle: "Search for a specific pizzeria"});

var pizzerias = $.ajax({
	dataType: "json",
	url: "pizzerias.geojson",
/*	success: function(data) {
/*		console.log(data);
/*	}, */
	error: function(xhr) {
		alert(xhr.statusText);
	}
})
$.when(pizzerias).done(function() {
	let pizzamap = L.map('mapid')
		.setView([51.505, -0.09], 2)
		.setMaxBounds([
			[-90, -180],
			[90, 180]
		]);

	let osm = L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
		maxZoom: 19,
		attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
	}).addTo(pizzamap);

	let dataLayer = new L.geoJSON(pizzerias.responseJSON, {
		onEachFeature: function(feature, layer) {
			feature.layer = layer;
		},
		pointToLayer: function(feature, latlng) {
			let pizzaIcon = L.Icon.extend({
				options: {
					iconUrl:	"slice_of_pizza.png",
					iconSize:		[32, 32],
					iconAnchor:		[16, 32]
				}
			});
			let address = feature.properties.address + ', ' + feature.properties.city;
			let name = feature.properties.pizzeria;
			let website = feature.properties.website;
			let pizzeriaPopup = L.popup({
				options: {
					autoclose: true
				}
			});
			return L.marker(latlng, {icon: new pizzaIcon()}).bindPopup('<b id="pizzerianame">' + name + '</b><br /><a href="' + website + '">' + website + '</a><br>' + address + '<br>');
		},
	}).addTo(pizzamap);

	// add search control to the map
	searchCtrl.indexFeatures(pizzerias.responseJSON, ['pizzeria', 'address', 'city']);
	searchCtrl.addTo(pizzamap);

	// ask to geolocate the user
	locateUser(pizzamap);
});
