;(function() {

    window._u = new function() {
        
        Object.defineProperty(
            this,
            "serverRoot",
            {
                get: function() {
                    return _createPath();
                }
            }
        );

        Object.defineProperty(
            this,
            "secureServerRoot",
            {
                get: function() {
                    return _createPath(true);
                }
            }
        );

        var _createPath = function(secureF) {
            var ll = window.location;
            var host = ll.host;
            var protocol = secureF ? "https:" : ll.protocol;
            var retval = protocol + "//" + host;
            return retval;
        };

        /**
         * Add comma for fee.
         */
        this.addComma = function( num ) {
            return String(num).replace( /(\d)(?=(\d\d\d)+(?!\d))/g, '$1,');
        };

        /**
         * Bind function and controll object scope.
         *
         * @param targetObj controlled scope
         * @param functionName function name.
         * @return wrapped function
         */
        this.bind = function( targetObj, functionName ) {
            var t = targetObj;
            var f = functionName;
            var retFunc = function() {
                return t[f].apply( t, arguments );
            };
            return retFunc;
        };

        /**
         * Check argument is String or not
         */
        this.isString = function( value ) {
            return (typeof(value) == "string") || (value instanceof String);
        };

        /**
         * Check argument is Array or not
         */
        this.isArray = function( value ) {
            return value instanceof Array;
        };

        /**
         * Check argument is empty or not
         *
         * @param value String | Array
         * @return true | false
         */
        this.isEmpty = function( value ) {

            var retval = true;
            if ( value != null ) {

                if ( this.isArray(value) && 0 < value.length ) {
                    retval = false;
                } else if ( this.isString(value) && 0 < value.length ) {
                    retval = false;
                }

            }

            return retval;
        };

        /**
         * Check argument has value or not
         *
         * @param value String | Array
         * @return true | false
         */
        this.hasValue = function( value ) {
            return !this.isEmpty( value );
        }

        /**
         * copy object
         */
        this.copyObject = function( dst, src ) {
            return $.extend( true, dst, src );
        };

        /**
         * Get int value, if value is string.
         *
         * @param value anything.
         * @return int value
         */
        this.getIntValue = function( value ) {

            var retval = null;
            if ( value instanceof String ) {
                retval = parseInt( value );
            } else if ( typeof value == 'number' ) {
                retval = value;
            } else {
                retval = 0;
            }

            return retval;
        };

        /**
         * Scroll to top
         */
        this.scrollToTop = function() {

            $('html,body').animate(
                { 'scrollTop' : 0 },
                1000,
                'swing'
            );

        };

    };

})(window);