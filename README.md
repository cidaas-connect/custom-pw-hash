# Deep Dive 4 - Smooth migration to cidaas

## Szenario

**Ausgangspunkt:** Bei der Migration werden die (hoffentlich) gehashten Passwörter aus dem Bestandssystem migriert. Da der verwendetet Hash-Algorithmus nicht Standardkonform ist, wird er von cidaas nicht unterstützt.

**Ziel:** Die Passwörter sollen trotzdem migriert werden. Beim ersten Login eines Nutzers soll das Passwort mit einem neuen, standardkonformen Hash-Algorithmus gehasht werden.


## Exercise - Erstelle deinen eigenen Passwort Hash Algorithmus

Lege über die [Postman Collection](https://www.postman.com/solar-water-27959/cidaas-connect-2024/overview) einen (oder mehrere) User mit einem frei erfundenen passwort hash an bspw. *meinhashprefixpasswort*.

Implementiere in der Datei *singlehash_controller.go* die Funktion *verifyCustomHashPassword* und prüfe ob das einkommende Klartext Password mit deinem Hash übereinstimmt.

Bsp.

```golang
func verifyCustomHashPassword(userPassword, storedHash string) bool {
    prefix := "meinhashprefix"
    return (prefix + userPassword) == storedHash
}
```

Push deinen code und versuche dich mit deinem User einzuloggen.

## Setup

Wir haben 5 Branches mit CI/CD vorbereitet

|Branchname|algorithmTypeId|
|-----|-----|
|1_custom_hash|CIDAAS_CONNECT_CUSTOM_HASH_1|
|2_custom_hash|CIDAAS_CONNECT_CUSTOM_HASH_2|
|3_custom_hash|CIDAAS_CONNECT_CUSTOM_HASH_3|
|4_custom_hash|CIDAAS_CONNECT_CUSTOM_HASH_4|
|5_custom_hash|CIDAAS_CONNECT_CUSTOM_HASH_5|