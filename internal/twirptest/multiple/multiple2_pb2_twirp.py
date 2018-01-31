# Code generated by protoc-gen-twirp_python v5.2.0, DO NOT EDIT.
# source: multiple2.proto

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

class Svc2Client(object):
    def __init__(self, server_address):
        """Creates a new client for the Svc2 service.

        Args:
            server_address: The address of the server to send requests to, in
                the full protocol://host:port form.
        """
        self.__target = server_address
        self.__service_name = "twirp.internal.twirptest.multiple.Svc2"

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

    def send(self, msg2):
        serialize = _sym_db.GetSymbol("twirp.internal.twirptest.multiple.Msg2").SerializeToString
        deserialize = _sym_db.GetSymbol("twirp.internal.twirptest.multiple.Msg2").FromString

        full_method = "/{}/{}".format(self.__service_name, "Send")
        body = serialize(msg2)
        resp_str = self.__make_request(body=body, full_method=full_method)
        return deserialize(resp_str)

    def same_package_proto_import(self, msg1):
        serialize = _sym_db.GetSymbol("twirp.internal.twirptest.multiple.Msg1").SerializeToString
        deserialize = _sym_db.GetSymbol("twirp.internal.twirptest.multiple.Msg1").FromString

        full_method = "/{}/{}".format(self.__service_name, "SamePackageProtoImport")
        body = serialize(msg1)
        resp_str = self.__make_request(body=body, full_method=full_method)
        return deserialize(resp_str)

