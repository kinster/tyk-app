// Add Channel Header Middleware (Apigee Add-Channel equivalent)
var addChannelHeader = new TykJS.TykMiddleware.NewMiddleware({});

addChannelHeader.NewProcessRequest(function(request, session, config) {
    var channel = "Web"; // default if not present on the key
    if (session && session.meta_data && session.meta_data.channel) {
        channel = session.meta_data.channel;
    }

    request.SetHeaders["X-CHANNEL"] = channel;
    request.SetHeaders["X-Forwarded-Prefix"] = "/dummy-proxy/v1";

    log("Added X-CHANNEL header: " + channel);
    return addChannelHeader.ReturnData(request, session && session.meta_data ? session.meta_data : {});
});

log("Add-channel middleware loaded");