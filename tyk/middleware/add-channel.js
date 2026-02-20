// Add Channel Header Middleware
// Equivalent to Apigee: Add-Channel Policy

var addChannelHeader = new TykJS.TykMiddleware.NewMiddleware({});

addChannelHeader.NewProcessRequest(function(request, session, config) {

    // Default channel
    var channel = "Web";

    // Read from API key metadata
    if (session && session.meta_data && session.meta_data.channel) {
        channel = session.meta_data.channel;
    }

    // Inject into headers
    request.SetHeaders["X-CHANNEL"] = channel;

    // Apigee-compat prefix header
    request.SetHeaders["X-Forwarded-Prefix"] = "/dummy-proxy/v1";

    log("Added X-CHANNEL header: " + channel);

    return addChannelHeader.ReturnData(request, session.meta_data);
});

log("Add-channel middleware loaded");