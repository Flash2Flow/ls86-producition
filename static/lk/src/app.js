$('#open-menu-nav').on('click', function(){
    $('#burger-menu').show();
    return false;
})
$('#person-block-1').on('click', function(){
  console.log("press")
  var name = $("#person-block-1").attr("name");
  window.location.replace("http://ls-86-rp.ru/home/" + name);
})
$('#menu_nav_mobile').on('click', function(){
  window.location.replace("http://ls-86-rp.ru/home/");
})
$('#menu_nav').on('click', function(){
  window.location.replace("http://ls-86-rp.ru/home/");
})
$('#person-block-2').on('click', function(){
  console.log("press")
  var name = $("#person-block-2").attr("name");
  window.location.replace("http://ls-86-rp.ru/home/" + name);
})

$('#person-block-3').on('click', function(){
  console.log("press")
  var name = $("#person-block-3").attr("name");
  window.location.replace("http://ls-86-rp.ru/home/" + name);
})
$('#person-block-4').on('click', function(){
  console.log("press")
  var name = $("#person-block-4").attr("name");
  window.location.replace("http://ls-86-rp.ru/home/" + name);
})



$('#btn-exit').on('click', function(){
  window.location.replace("http://ls-86-rp.ru/exit");
})
$('#btn-settings').on('click', function(){
    $('#block-settings-account').show(300);
    return false;
})
$('#btn-create-pers').on('click', function(){
  window.location.replace("http://ls-86-rp.ru/ucp");
  return false;
})
$(document).ready(function() {
  $('#person-block').show(300);
});
$('#pencil').on('click', function(){
    $('#block-remove-password').show(300);
    $('#block-settings-account').hide();
    return false;
})
$('#btn-accept').on('click', function(){
    $('#block-set-password').show(300);
    $('#block-remove-password').hide();
    $('#block-settings-account').hide();
    return false;
})

$(document).click( function(event){
    if( $(event.target).closest("#burger-menu").length ) return;
    $("#burger-menu").hide();
    event.stopPropagation();
  });

  $(document).click( function(event){
    if( $(event.target).closest("#block-settings-account").length ) return;
    $("#block-settings-account").hide(300);
    event.stopPropagation();
  });

  $(document).click( function(event){
    if( $(event.target).closest("#block-remove-password").length ) return;
    $("#block-remove-password").hide(300);
    event.stopPropagation();
  });
  $(document).click( function(event){
    if( $(event.target).closest("#block-set-password").length ) return;
    $("#block-set-password").hide(300);
    event.stopPropagation();
  });