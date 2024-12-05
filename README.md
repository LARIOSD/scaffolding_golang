# Scaffolding para Arquitectura Hexagonal en Go (Golang)  

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)  

Este proyecto es un **scaffolding** diseñado para facilitar la creación de aplicaciones en Go utilizando la **Arquitectura Hexagonal**. Proporciona una base sólida y modular para desarrollar aplicaciones escalables, mantenibles y adaptables a diferentes infraestructuras tecnológicas.  

## Características Principales  
- **Estructura Modular**: Separación clara en capas (`domain`, `infrastructure`, `application`, `presentation`) para fomentar el desacoplamiento y facilitar las pruebas.  
- **Inyección de Dependencias**: Simplifica la gestión de dependencias, permitiendo mayor flexibilidad en la integración de implementaciones.  
- **Soporte Multi-Base de Datos**: Ejemplo de integración con PostgreSQL, MongoDB, y otros motores, manteniendo independencia tecnológica.  
- **Rutas y Controladores**: Implementación de controladores y rutas RESTful como punto de inicio para desarrollar APIs robustas.  
- **Pruebas Unitarias y de Integración**: Estructura preconfigurada para agregar pruebas fácilmente.  
- **Ejemplo Funcional**: Contiene ejemplos básicos para operaciones CRUD.  
- **Flexibilidad y Extensibilidad**: Preparado para adaptarse a diferentes dominios y casos de uso.  

## Estructura del Proyecto  
```plaintext
├── application       # Casos de uso y servicios de la aplicación
├── domain            # Entidades y lógica del negocio
├── infrastructure    # Implementaciones técnicas (DB, APIs externas)
├── presentation      # Rutas, controladores y transporte HTTP
└── main.go           # Punto de entrada de la aplicación
