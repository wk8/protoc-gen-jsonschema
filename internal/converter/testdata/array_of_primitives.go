package testdata

const ArrayOfPrimitives = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "properties": {},
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "string"
                }
            ]
        },
        "luckyNumbers": {
            "items": {
                "oneOf": [
                    {
                        "type": "null"
                    },
                    {
                        "type": "integer"
                    }
                ]
            },
            "properties": {},
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        },
        "luckyBigNumbers": {
            "items": {
                "oneOf": [
                    {
                        "type": "integer"
                    },
                    {
                        "type": "string"
                    },
                    {
                        "type": "null"
                    }
                ]
            },
            "properties": {},
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        },
        "keyWords": {
            "items": {
                "oneOf": [
                    {
                        "type": "null"
                    },
                    {
                        "type": "string"
                    }
                ]
            },
            "properties": {},
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        },
        "big_number": {
            "properties": {},
            "oneOf": [
                {
                    "type": "integer"
                },
                {
                    "type": "string"
                },
                {
                    "type": "null"
                }
            ]
        }
    },
    "additionalProperties": true,
    "oneOf": [
        {
            "type": "null"
        },
        {
            "type": "object"
        }
    ]
}`


const ArrayOfPrimitivesDouble = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "properties": {},
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "string"
                }
            ]
        },
        "luckyNumbers": {
            "items": {
                "oneOf": [
                    {
                        "type": "null"
                    },
                    {
                        "type": "integer"
                    }
                ]
            },
            "properties": {},
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        },
        "luckyBigNumbers": {
            "items": {
                "oneOf": [
                    {
                        "type": "integer"
                    },
                    {
                        "type": "string"
                    },
                    {
                        "type": "null"
                    }
                ]
            },
            "properties": {},
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        },
        "keyWords": {
            "items": {
                "oneOf": [
                    {
                        "type": "null"
                    },
                    {
                        "type": "string"
                    }
                ]
            },
            "properties": {},
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        },
        "big_number": {
            "properties": {},
            "oneOf": [
                {
                    "type": "integer"
                },
                {
                    "type": "string"
                },
                {
                    "type": "null"
                }
            ]
        },
        "bigNumber": {
            "properties": {},
            "oneOf": [
                {
                    "type": "integer"
                },
                {
                    "type": "string"
                },
                {
                    "type": "null"
                }
            ]
        }
    },
    "additionalProperties": true,
    "oneOf": [
        {
            "type": "null"
        },
        {
            "type": "object"
        }
    ]
}`

