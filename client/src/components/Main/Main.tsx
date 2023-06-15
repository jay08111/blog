import RecentPosts from '../RecentPosts/RecentPosts'
import s from './Main.module.css'

const Main: React.FC = () => {
   return (
      <div className={s.mainContainer}>
         <main>
            <RecentPosts />
         </main>
      </div>
   )
}

export default Main
