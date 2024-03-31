// Using Javascript bc i don't know how to do it in htmx and golang yet :/

console.log('This is the ViewTickets.js file')
const modal = document.getElementById('modal')
const openModal = document.getElementById('new-ticket-btn')
const closeModal = document.getElementById('close-modal')

openModal.addEventListener('click', () => {
  modal.showModal()
})

closeModal.addEventListener('click', () => {
  modal.close()
})
