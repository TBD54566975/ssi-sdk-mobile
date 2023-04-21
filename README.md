### ⚠️ Please note that this package is a prototype and is not intended for production. Breaking changes will appear without warning. Future adoption is uncertain.

<br /> <br />

# Monorepo for TBD's SSI SDK

## Layout

| Resource                               | Description                                                              |
| -------------------------------------- | ------------------------------------------------------------------------ |
| [sdk](./sdk)                           | Business logic core written in Go and outputted to xcframeworks and AARs |
| [react-native-ssi](./react-native-ssi) | React Native SDK wrapper overtop of the business logic core              |
| [package.json](./package.json)         | Used for monorepo automations together with lefthook.yml                 |
