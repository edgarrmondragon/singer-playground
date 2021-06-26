import singer


ORDERS_SCHEMA = {
    "properties": {
        "id": {"type": "integer"},
        "nested": {
            "type": "object",
            "properties": {"a": {"type": "integer"}, "b": {"type": "boolean"}},
        },
    }
}

schemas = {
    "123_orders": ORDERS_SCHEMA,
    "456_orders": ORDERS_SCHEMA,
}

records = {
    "123_orders": [
        {"id": 1, "nested": {"a": 42, "b": False}},
        {"id": 2, "nested": {"a": 314, "b": True}},
    ],
    "456_orders": [
        {"id": 1, "nested": {"a": 1513, "b": False}},
    ],
}

state = {"123_orders": 2, "456_orders": 1}

for stream, schema in schemas.items():
    singer.write_schema(stream, schema, key_properties=["id"])

for stream, record in records.items():
    singer.write_records(stream, record)

singer.write_state(state)
