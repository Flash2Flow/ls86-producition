$('#open-menu-nav').on('click', function(){
    $('#burger-menu').show();
    return false;
})
$('#btn-settings').on('click', function(){
    $('#block-settings-account').show(300);
    return false;
})
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