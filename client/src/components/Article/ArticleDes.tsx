import React from 'react'
import Badge from '../../view/base/Badge'

type Iprops = {
   updated_at?: string
   writer?: string
   title: string
   content: string
   badge: string[]
}

const ArticleDes = (props: Iprops) => {
   const {
      updated_at = new Date(),
      writer = 'hoyeoun',
      title,
      content,
      badge,
   } = props

   return (
      <article>
         <div>
            <p>{updated_at.toString() ?? ''}</p>
            <p>{writer ?? ''}</p>
         </div>
         <div>
            <h3>{title ?? ''}</h3>
            <p>{content ?? ''}</p>
         </div>
         <Badge badge={badge} />
      </article>
   )
}

export default ArticleDes
