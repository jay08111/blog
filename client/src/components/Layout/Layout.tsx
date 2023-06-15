import Header from '../Header/Header'
import s from './Layout.module.css'
import Main from '../Main/Main'

interface Props {
   children?: React.ReactNode
}

const Layout: React.FC<Props> = ({ children }) => {
   return (
      <>
         <div className={s.allWrapper}>
            <Header />
            <main>{children}</main>
         </div>
      </>
   )
}

export default Layout
