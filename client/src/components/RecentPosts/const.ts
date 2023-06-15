type RECENT_POST_TYPE = {
    id: number
    category?: number
    sub_category?: number
    title: string
    img: string
    content: string
    writer?: string
 }
 
 const MOCK_RECENT_POST: RECENT_POST_TYPE[] = [
    {
       id: 1,
       category: 1,
       sub_category: 1,
       title: 'Ux review presentations',
       img: 'https://images.unsplash.com/photo-1682687220640-9d3b11ca30e5?ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=3540&q=80',
       content:
          'How do you  create compelling presentations that wow your colleagues and impress your managers?',
    },
    {
       id: 2,
       category: 1,
       sub_category: 1,
       title: 'Ux review presentations',
       img: 'https://images.unsplash.com/photo-1682687220640-9d3b11ca30e5?ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=3540&q=80',
       content:
          'How do you  create compelling presentations that wow your colleagues and impress your managers?',
    },
    {
       id: 3,
       category: 1,
       sub_category: 1,
       title: 'Ux review presentations',
       img: 'https://images.unsplash.com/photo-1682687220640-9d3b11ca30e5?ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=3540&q=80',
       content:
          'How do you  create compelling presentations that wow your colleagues and impress your managers?',
    },
    {
       id: 4,
       category: 1,
       sub_category: 1,
       title: 'Ux review presentations',
       img: 'https://images.unsplash.com/photo-1682687220640-9d3b11ca30e5?ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=3540&q=80',
       content:
          'How do you  create compelling presentations that wow your colleagues and impress your managers?',
    },
 ]
 
 export { MOCK_RECENT_POST }
 