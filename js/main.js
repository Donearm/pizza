let mymap = L.map('mapid').setView([51.505, -0.09], 2);
mymap.setMaxBounds([
	[-90, -180],
	[90, 180]
]);

let osm = L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
	maxZoom: 19,
	attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
}).addTo(mymap);

let xhr = new XMLHttpRequest();
xhr.open('GET', "pizzerias.geojson");
xhr.setRequestHeader('Content-Type', 'application/json');
xhr.responseType = 'json';
xhr.onload = function() {
	if (xhr.status !== 200) { return }
	let dataLayer = L.geoJSON(xhr.response, {
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
			return L.marker(latlng, {icon: new pizzaIcon()}).bindPopup('<b>' + name + '</b><br /><a href="' + website + '">' + website + '</a><br>' + address + '<br>');
		},
	});
	dataLayer.addTo(mymap);
};
xhr.send();
