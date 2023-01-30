package human.resource.mgmt.infra;

import human.resource.mgmt.domain.*;
import org.springframework.hateoas.EntityModel;
import org.springframework.hateoas.Link;
import org.springframework.hateoas.server.RepresentationModelProcessor;
import org.springframework.stereotype.Component;

@Component
public class CalendarHateoasProcessor
    implements RepresentationModelProcessor<EntityModel<Calendar>> {

    @Override
    public EntityModel<Calendar> process(EntityModel<Calendar> model) {
        model.add(
            Link
                .of(model.getRequiredLink("self").getHref() + "/add")
                .withRel("add")
        );
        model.add(
            Link
                .of(model.getRequiredLink("self").getHref() + "/cancel")
                .withRel("cancel")
        );
        model.add(
            Link
                .of(model.getRequiredLink("self").getHref() + "/delayschedule")
                .withRel("delayschedule")
        );

        return model;
    }
}
