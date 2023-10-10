## Difference Between Map and Struct
* maps are passed by reference by default but when passing struct it is passsed by value.
* To delete a key in map ,we do like this:
delete(map_name,'key_name')

Certainly, here's a comparison between structs and maps in Go in a tabular form:

| Aspect                  | Structs                                   | Maps                                       |
|-------------------------|------------------------------------------|--------------------------------------------|
| Purpose                 | Used to define custom data types with named fields, making it suitable for representing structured data. | Used to store key-value pairs, allowing for dynamic data association and retrieval. |
| Declaration Syntax      | Defined using the `type` keyword and field names with their data types. | Created using the `make` function or a composite literal with key-value pairs. |
| Example                 | ```go                                       | ```go                                       |
|                         | type Person struct {                        | person := map[string]string{              |
|                         |     FirstName string                        |     "FirstName": "John",                  |
|                         |     LastName  string                        |     "LastName":  "Doe",                   |
|                         | }                                          | }                                          |
| Accessing Fields        | Access fields using dot notation: `person.FirstName` | Access values by providing the key: `name := person["FirstName"]` |
| Use Case                | Suitable for structured data with known fields and fixed structure, like representing a person's attributes. | Useful for dynamic data where keys and values can change during runtime, like storing configuration settings or dynamic data aggregation. |
| Type Safety             | Provides strong type safety because field types are defined at compile time. | Key types can vary and are often `string`, which may lead to runtime errors if not handled correctly. |
| Fixed Structure         | Has a fixed structure defined by the fields, making it less flexible for dynamic data. | Offers flexibility because keys and values can be added or removed easily. |
| Performance             | Generally more efficient for static, structured data due to compile-time optimizations. | May be less efficient for very large datasets or when frequent dynamic modifications are required. |
| Initialization          | Requires specifying values for all fields when creating an instance. | Can be created and updated incrementally, adding keys and values as needed. |
| Memory Allocation       | Structs have a fixed memory layout, which can be advantageous for memory efficiency. | Maps may have some overhead due to dynamic resizing and hashing. |
| Comparison              | Structs are typically compared by comparing each field individually. | Maps are compared by checking if they reference the same underlying map or if they are both nil. |
| Examples                | Representing a car with attributes like make, model, and year. | Storing key-value pairs for a phone book, where names are keys and phone numbers are values. |

In summary, structs are suitable for defining custom data structures with a fixed layout, while maps are used for dynamic key-value storage with flexibility in adding and removing entries. The choice between them depends on the specific requirements of your data and how you plan to use it in your Go program.