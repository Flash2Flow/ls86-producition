$('#active_block_join_one').on('click', function(){
    $('#block_join').show();
    $('#close-menu-auth').show();
    $('#burger-menu').hide();
    return false;
})
$('#close-menu-auth').on('click', function(){
  $('#block_join').hide();
  return false;
})
$('#close-rule-server').on('click', function(){
  $('#up-block-rule-server').hide();
  return false;
})
$('#click_rule').on('click', function(){
  $('#up-block-rule-server').show();
  $('#up-block-rule-join').hide();
  $('#up-block-rule-created').hide();
  return false;
})
$('#click_create').on('click', function(){
  $('#up-block-rule-created').show();
  $('#up-block-rule-join').hide();
  $('#up-block-rule-server').hide();
  return false;
})
$('#click_join').on('click', function(){
  $('#up-block-rule-join').show();
  $('#up-block-rule-created').hide();
  $('#up-block-rule-server').hide();
  return false;
})
$('#close-created-server').on('click', function(){
  $('#up-block-rule-created').hide();
  return false;
})
$('#close-join-server').on('click', function(){
  $('#up-block-rule-join').hide();
  return false;
})
$('#close-menu').on('click', function(){
  $('#burger-menu').hide();
  return false;
})
$('#active_block_join_two').on('click', function(){
  $('#block_join').show();
  return false;
})
$('#burger-button').on('click', function(){
  $('#burger-menu').show();
  return false;
})
$('#reg').on('click', function(){
  $('#block_created').show();
  $('#block_join').hide();
  return false;
})