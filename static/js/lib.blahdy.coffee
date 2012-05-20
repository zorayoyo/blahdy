class BlahdyClient
  constructor: (@network) ->
    @apiBase = 'http://127.0.0.1:8888/api/'

  defaultErrorHandler: (url, xhr, textStatus, errorThrown)->
    console.log('Error:', url, xhr, textStatus, errorThrown)

  get: (ajax_url, ajax_params, on_success, on_error) ->
    @do_ajax('GET', ajax_url, ajax_params, {},
      (result, textStatus, xhr)->
        #@success_handler(result, textStatus, xhr)
        on_success(result, textStatus, xhr)
      ,
      (xhr, textStatus, errorThrown)->
        if on_error == undefined or on_error == null
          @defaultErrorHandler(ajax_url, xhr, textStatus, errorThrown)
        else
          on_error(xhr, textStatus, errorThrown)
    )

  post: (ajax_url, ajax_params, on_success, on_error)->
    @do_ajax('POST', ajax_url, ajax_params, {},
      (result, textStatus, xhr)->
        #@success_handler(result, textStatus, xhr)
        on_success(result, textStatus, xhr)
      ,
      (xhr, textStatus, errorThrown)->
        if on_error == undefined or on_error == null
          @defaultErrorHandler(ajax_url, xhr, textStatus, errorThrown)
        else
          on_error(xhr, textStatus, errorThrown)
    )

  do_ajax: (method, url, params, headers, on_success, on_error)->
    if method == 'GET'
      arr = []
      for k, v of params
        arr.push(encodeURIComponent(k) + '=' + encodeURIComponent(v))
      url = url + '?' + arr.join('&')
      params = {}
    # @TODO Authorization
    # headers['Authorization'] = self.basic_auth()
    @network.do_request(method, url, params, headers, null, on_success, on_error)

  getBlahList: (on_success)->
    url = @apiBase + 'blah/all'
    params = {}
    @get(url, params, on_success)

root = exports ? this
root.BlahdyClient = root.BlahdyClient ? BlahdyClient



