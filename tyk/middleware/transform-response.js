// Transform Response Middleware
// Runs after the upstream backend responds (response hook)
var transformResponse = new TykJS.TykMiddleware.NewMiddleware({});

transformResponse.NewProcessRequest(function(request, session, config) {
    // Response body is available as a base64-encoded string in request.Body
    var rawBody = b64dec(request.Body);

    try {
        var body = JSON.parse(rawBody);

        // 1. Rename "message" to "status"
        if (body.message !== undefined) {
            body.status = body.message;
            delete body.message;
        }

        // 2. Rename "channelReceived" to "channel"
        if (body.channelReceived !== undefined) {
            body.channel = body.channelReceived;
            delete body.channelReceived;
        }

        // 3. Rename "prefixReceived" to "prefix"
        if (body.prefixReceived !== undefined) {
            body.prefix = body.prefixReceived;
            delete body.prefixReceived;
        }

        // 4. Add metadata
        body.transformed = true;

        // Re-encode body back to base64
        request.Body = b64enc(JSON.stringify(body));

        log("transform-response: body transformed successfully");

    } catch (e) {
        log("transform-response: failed to parse body - " + e);
    }

    return transformResponse.ReturnData(request, {});
});

log("Transform-response middleware loaded");
