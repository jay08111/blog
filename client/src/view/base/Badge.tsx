import React from 'react'

type Iprops = {
   badge: string[]
}

function Badge(props: Iprops) {
   const { badge } = props
   return <div>{badge}</div>
}

export default Badge
