import React, {useContext} from "react";

import { DataContext } from "../../context/Dataprovider"
import { ProductoItem } from "./ProductoItem"

export const ProductosLista = () => {

        const value = useContext(DataContext)
        const [productos] = value.productos;

        console.log(productos)

    return (
        <>
        <h1 className="title">PRODUCTOS</h1>
        <div className="productos">
           {
               productos.map(producto =>(
                < ProductoItem 
                key={producto.id}
                title={producto.title}
                image={producto.image}
                category={producto.category}
                price={producto.price}
                id={producto.id}
                
                />
               ))
           }
        </div>
        </>
    )
}
