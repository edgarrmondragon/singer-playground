import sys

import singer


for message_string in iter(sys.stdin.readline, ""):
    message = singer.parse_message(message_string)

    if hasattr(message, "stream"):
        schema, table = message.stream.split("_", 1)
        message.stream = table

    if isinstance(message, singer.RecordMessage):
        message.record["__partition"] = schema
        singer.write_message(message)

    elif isinstance(message, singer.SchemaMessage):
        message.schema["properties"]["__partition"] = {"type": "string"}
        message.key_properties.append("__partition")
        singer.write_message(message)

    else:
        singer.write_message(message)
