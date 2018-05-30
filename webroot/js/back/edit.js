$(function () {

    window.ID_TEXTAREA = 1;
    window.ID_TEXTINPUT = 2;
    window.ID_TEXT_DROPDOWN = 3;
    window.ID_CHECKBOX = 4;
    window.ID_RADIO_BUTTON = 5;
    window.ID_HIDDEN = 6;
    window.ID_CONTENTS_SEQUENCE = 7;
    
    /** 
     * Get editing data from id arrays.
     * 
     * @param checkIDs id arrays
     * @return editing data object
     */
    window.getEditingData = function( checkIDs ) {

        var retval = null;
        for ( var kk in checkIDs ) {

            var vv = checkIDs[ kk ];
            var changedValue = getComponentValue( kk, vv );
            if ( changedValue != null ) {
                retval = retval || {};
                retval[ kk ] = changedValue;
            }

        }

        return retval;
    };

    /** 
     * Get value from jquery component
     * 
     * @param key ID name
     * @param value component type
     */
    window.getComponentValue = function( key, value ) {
        
        var retval = null;
        var $element = $( '#' + key );
        switch ( value ) {
            case ID_TEXTAREA:
            case ID_TEXTINPUT: {
                // textinput, textarea
                retval = $element.val();
                break;
            }
            case ID_TEXT_DROPDOWN: {
                // Dropdown
                retval = $element.find( 'option:selected' ).val();
                break;
            }
            case ID_CHECKBOX: {
                // Checkbox
                retval = $element.prop('checked');
                break;
            }
            case ID_HIDDEN: {
                // Hidden
                retval = parseInt($element.val());
                break;
            }
            case ID_CONTENTS_SEQUENCE: {
                // Contents Sequence
                var items = [];
                $li = $element.find('li');
                $li.each(function( ii, elem ) {
                    $ee = $(elem);
                    items.push( $ee.data('value') );
                });
                retval = items.join('_');
                break;
            }
            default: {
                break;
            }
        }

        return retval;
    };

    /**
     * Show baloon as tooltip
     * 
     * @param message string
     * @param callback event handler
     */
    window.showBalooon = function( message, callback ) {
        $('#errorMessage').html( message );
        $('#ballon').fadeIn(500).delay(1000).fadeOut(1000, callback);
    };

});