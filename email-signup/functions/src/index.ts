import * as functions from 'firebase-functions';
//
// The Firebase Admin SDK to access Cloud Firestore.
import * as admin from 'firebase-admin';
admin.initializeApp();

// // Start writing Firebase Functions
// // https://firebase.google.com/docs/functions/typescript
//
export const insertLead = functions.https.onRequest(async (request, response) => {
    if (request.method !== "POST") {
        response.status(400).send("WRONG METHOD");
        return;
    }
    if (request.body["email"] === undefined) {
        response.json({
            err: "invalid",
        });
        response.redirect("https://openschool.dev?newsletter=false");
        return;
    }
    await admin.firestore().collection("leads").add({
        email: request.body.email,
    });
    response.redirect("https://openschool.dev?newsletter=true");
});
