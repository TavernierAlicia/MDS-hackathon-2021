
function DisplayMap(lat = 48.9548585, lon = 2.8857861) {
  var mymap = L.map('my_osm_widget_map', { 
    center: [lat, lon],
    zoom: 17, 
    zoomControl: false, 
    scrollWheelZoom: false 
  });

  L.tileLayer('https://api.mapbox.com/styles/v1/aliciatavernier/ckl9kjckm10qp17o0s8389twc.html?fresh=true&title=view&access_token=pk.eyJ1IjoiYWxpY2lhdGF2ZXJuaWVyIiwiYSI6ImNrbDlrMDJmbDBhZ3Iybm9vcXBncnc4NHgifQ.8z1DDpgg5DtBuHvX-MRUmw', { 
    maxZoom: 20, 
    attribution: 'Données &copy; Contributeurs <a href="http://openstreetmap.org">OpenStreetMap</a> + ' +
      '<a href="http://mapbox.com">Mapbox</a> | ' +
      '<a href="https://creativecommons.org/licenses/by/2.0/">CC-BY</a> ' +
      'Guillaume Rouan 2016', 
  }).addTo(mymap);

  L.marker([lat, lon]).addTo(mymap);
}