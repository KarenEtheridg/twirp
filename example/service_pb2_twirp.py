# Code generated by protoc-gen-twirp_python v5.1.0, DO NOT EDIT.
# source: service.proto

import httplib
import json
import urllib2
from google.protobuf import symbol_database as _symbol_database

_sym_db = _symbol_database.Default()

class TwirpException(httplib.HTTPException):
    def __init__(self, code, message, meta):
        self.code = code
        self.message = message
        self.meta = meta

    @classmethod
    def from_http_err(cls, err):
        try:
            jsonerr = json.load(err)
            code = jsonerr["code"]
            msg = jsonerr["msg"]
            meta = jsonerr.get("meta")
            if meta is None:
                meta = {}
        except:
            code = "internal"
            msg = "Error from intermediary with HTTP status code {} {}".format(
                err.code, httplib.responses[err.code],
            )
            meta = {}
        return cls(code, msg, meta)

class HaberdasherClient(object):
    """
    A Haberdasher makes hats for clients.
    """

    def __init__(self, server_address):
        """Creates a new client for the Haberdasher service.

        Args:
            server_address: The address of the server to send requests to, in
                the full protocol://host:port form.
        """
        self.__target = server_address
        self.__service_name = "twitch.twirp.example.Haberdasher"

    def __make_request(self, body, full_method):
        req = urllib2.Request(
            url=self.__target + "/twirp" + full_method,
            data=body,
            headers={"Content-Type": "application/protobuf"},
        )
        try:
            resp = urllib2.urlopen(req)
        except urllib2.HTTPError as err:
            raise TwirpException.from_http_err(err)

        return resp.read()

    def make_hat(self, size):
        """
        MakeHat produces a hat of mysterious, randomly-selected color!
        """

        serialize = _sym_db.GetSymbol("twitch.twirp.example.Size").SerializeToString
        deserialize = _sym_db.GetSymbol("twitch.twirp.example.Hat").FromString

        full_method = "/{}/{}".format(self.__service_name, "MakeHat")
        body = serialize(size)
        resp_str = self.__make_request(body=body, full_method=full_method)
        return deserialize(resp_str)

