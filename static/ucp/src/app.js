$('#open-menu-nav').on('click', function(){
    $('#burger-menu').show();
    return false;
})
$('#menu_nav_mobile').on('click', function(){
    window.location.replace("http://localhost:3030/home/");
})
$('#menu_nav').on('click', function(){
    window.location.replace("http://localhost:3030/home/");
})


$(document).click( function(event){
    if( $(event.target).closest("#burger-menu").length ) return;
    $("#burger-menu").hide();
    event.stopPropagation();
  });