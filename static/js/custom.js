$(document).ready(function(e) {
    jQuery(function(){
        jQuery('.showStage1').click(function(){
            jQuery('.card').hide();
            jQuery('.targetStage1').show();
        });
        jQuery('.showStage2').click(function(){
            jQuery('.card').hide();
            jQuery('.targetStage2').show();
        });
        jQuery('.showStage2').click(function(){
            jQuery('.card').hide();
            jQuery('.targetStage2').show();
        });
        jQuery('.showStage3').click(function(){
            jQuery('.card').hide();
            jQuery('.targetStage3').show();
        });
        jQuery('.showStage4').click(function(){
            jQuery('.card').hide();
            jQuery('.targetStage4').show();
        });
        jQuery('.showStage5').click(function(){
            jQuery('.card').hide();
            jQuery('.targetStage5').show();
        });
        jQuery('.showStage6').click(function(){
            jQuery('.card').hide();
            jQuery('.targetStage6').show();
        });
        $('img[usemap]').rwdImageMaps();
    });
});


function changeCard(quest, np){
    var maxquest  = 10;
    var parentpro = document.getElementById("pro");
    var childpro  = document.getElementById("probar");
    var cards     = document.getElementsByClassName("calc-card");
    if (np == "next") quest++;
    else if (np == "prev") quest--;

    for (var i = 0; i <cards.length; i++ ){
        cards[i].style.display = "none";
    }
    var newCard = document.getElementById("Question"+quest);
    newCard.style.display = "block"; 
    if(quest == 11){
        parentpro.style.display = "none";
    } else {
        parentpro.style.display = "block";
        w = quest / maxquest * 100;
        childpro.innerText = quest + "/" + maxquest; 
        childpro.style.width = w + "%";
    }

}    

function sticky_close(){
    var sticky = document.getElementById("sticky");
    sticky.style.display = "none";
}


function submit(){

}
