;(function( window ) {
    
    window._s = new function() {

        //-------------------------------------------------------------------------
        // Consts
        //-------------------------------------------------------------------------
        // response:success
        this.LEVEL_SUCCESS = 0;
        // resuponse:error
        this.LEVEL_ERROR = -1;

        // Content Type
        this.CONTENT_TYPE_JSON = 'application/JSON';

        // response data type
        this.DATATYPE_TEXT = 'text';
        // response json type
        this.DATATYPE_JSON = 'json';
        // response html type
        this.DATATYPE_HTML = 'html';

        // method, get
        this.METHOD_GET = 'GET';
        // method, post
        this.METHOD_POST = 'POST';

        //-------------------------------------------------------------------------
        // Property
        //-------------------------------------------------------------------------
        // requesting flag, block multi request
        var _isRequesting = false;

        // suspend param buffer
        var _requestQueue = [];

        //-------------------------------------------------------------------------
        // public method.
        //-------------------------------------------------------------------------
        /**
         * Call service of concrete function
         *
         * @param serviceURL service url
         * @param param body parameter
         * @param call back function
         */
        this.callService = function( serviceURL, resultHandler, option ) {

            if ( _isRequesting ) {
                _addQueue( serviceURL, resultHandler, option );
                return;
            }

            var that = this;
            var handler = resultHandler;
            var method = (option != null && option.method != null) ? option.method : this.METHOD_GET;
            var param = (method == this.METHOD_POST) ? option.param : null;
            var dataType = (option != null && option.dataType != null) ? option.dataType : null;
            var contentType = (option != null && option.contentType != null) ? option.contentType : null;
            
            var inquiry = {
                url: serviceURL,
                data: param,
                type: method,
                dataType: dataType,
                contentType: contentType,
                scriptCharset: 'utf-8',
                cache: false,
                asynch: true,
                success: function( data, status, xhr ) {
                    that.successHandler( handler, data, status );
                },
                error: function( xhr, textStatus, errorThrown ) {
                    that.errorHandler( handler, xhr, textStatus, errorThrown );
                },
                complete: function( xhr, textStatus ) {
                    console.log( "Complete : " + serviceURL );
                }
            };

            // send request
            $.ajax( inquiry );

            _isRequesting = true;
        };

        //-------------------------------------------------------------------------
        // private
        //-------------------------------------------------------------------------
        /**
         * add queue
         */
        var _addQueue = function( serviceURL, resultHandler, option ) {

            var tmp = {};
            tmp.url = serviceURL;
            tmp.handler = resultHandler ;
            tmp.option = option ;

            _requestQueue.push( tmp );
        };

        /**
         * get queue
         */
        var _getQueue = function() {

            var retval = null;
            if ( 0 < _requestQueue.length ) {
                // reverse sequence and get the tail object
                retval = _requestQueue.reverse().pop();
                // reverse sequence
                _requestQueue.reverse();
            }

            return retval;
        };

        /**
         *
         */
        var _checkQueue = function() {

            var queue = _getQueue();
            if ( queue != null ) {

                window._s.callService(
                    queue.url,
                    queue.handler,
                    queue.option
                );

            }

        };

        //-------------------------------------------------------------------------
        // Service handler
        //-------------------------------------------------------------------------
        /**
         * Success handler
         */
        this.successHandler = function( handler, data, textStatus ) {

            _isRequesting = false;
            _checkQueue();

            handler( window._s.LEVEL_SUCCESS, data );
        };

        /**
         * Error handler.
         */
        this.errorHandler = function( handler, xhr, textStatus, errorThrown ) {

            _isRequesting = false;
            _checkQueue();

            var data = xhr.responseText;
            handler( window._s.LEVEL_ERROR, data );
        };

    };

    /**
     * Call service
     *
     * @param path service path
     * @param
     */
    window.callService = function( path, param, callback ) {

        var jsonParam = JSON.stringify(param);
        var option = {
            'param' : jsonParam,
            'method' : _s.METHOD_POST,
            'dataType': _s.DATATYPE_JSON,
            'contentType' : _s.CONTENT_TYPE_JSON,
        };

        window._s.callService(
            _u.serverRoot + path,
            callback,
            option
        );

    };

}) ( window );
    